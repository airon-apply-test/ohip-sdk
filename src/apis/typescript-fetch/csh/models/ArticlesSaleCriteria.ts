/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Cashiering API
 * APIs to cater for Cashiering related functionality in OPERA Cloud.<br /><br /> Cashiering provides access to a guest folio, posting journals, receipt histories, currency calculations, credit card settlements, and check a guest out.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import type { InstanceLink } from './InstanceLink';
import {
    InstanceLinkFromJSON,
    InstanceLinkFromJSONTyped,
    InstanceLinkToJSON,
} from './InstanceLink';
import type { SaleCriteriaType } from './SaleCriteriaType';
import {
    SaleCriteriaTypeFromJSON,
    SaleCriteriaTypeFromJSONTyped,
    SaleCriteriaTypeToJSON,
} from './SaleCriteriaType';
import type { WarningType } from './WarningType';
import {
    WarningTypeFromJSON,
    WarningTypeFromJSONTyped,
    WarningTypeToJSON,
} from './WarningType';

/**
 * Request to post a set of charges, payments and generate a bill for a PasserBy or for the Post-It functionality . A PasserBy is someone who is not staying at the Property. A PasserBy transaction requires a Profile in Opera for whom the the folio is being created.
 * @export
 * @interface ArticlesSaleCriteria
 */
export interface ArticlesSaleCriteria {
    /**
     * 
     * @type {SaleCriteriaType}
     * @memberof ArticlesSaleCriteria
     */
    criteria?: SaleCriteriaType;
    /**
     * 
     * @type {Array<InstanceLink>}
     * @memberof ArticlesSaleCriteria
     */
    links?: Array<InstanceLink>;
    /**
     * Used in conjunction with the Success element to define a business error.
     * @type {Array<WarningType>}
     * @memberof ArticlesSaleCriteria
     */
    warnings?: Array<WarningType>;
}

/**
 * Check if a given object implements the ArticlesSaleCriteria interface.
 */
export function instanceOfArticlesSaleCriteria(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function ArticlesSaleCriteriaFromJSON(json: any): ArticlesSaleCriteria {
    return ArticlesSaleCriteriaFromJSONTyped(json, false);
}

export function ArticlesSaleCriteriaFromJSONTyped(json: any, ignoreDiscriminator: boolean): ArticlesSaleCriteria {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'criteria': !exists(json, 'criteria') ? undefined : SaleCriteriaTypeFromJSON(json['criteria']),
        'links': !exists(json, 'links') ? undefined : ((json['links'] as Array<any>).map(InstanceLinkFromJSON)),
        'warnings': !exists(json, 'warnings') ? undefined : ((json['warnings'] as Array<any>).map(WarningTypeFromJSON)),
    };
}

export function ArticlesSaleCriteriaToJSON(value?: ArticlesSaleCriteria | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'criteria': SaleCriteriaTypeToJSON(value.criteria),
        'links': value.links === undefined ? undefined : ((value.links as Array<any>).map(InstanceLinkToJSON)),
        'warnings': value.warnings === undefined ? undefined : ((value.warnings as Array<any>).map(WarningTypeToJSON)),
    };
}

