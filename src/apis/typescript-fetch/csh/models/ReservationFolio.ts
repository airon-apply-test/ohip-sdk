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
import type { CreateProformaCriteriaType } from './CreateProformaCriteriaType';
import {
    CreateProformaCriteriaTypeFromJSON,
    CreateProformaCriteriaTypeFromJSONTyped,
    CreateProformaCriteriaTypeToJSON,
} from './CreateProformaCriteriaType';
import type { InstanceLink } from './InstanceLink';
import {
    InstanceLinkFromJSON,
    InstanceLinkFromJSONTyped,
    InstanceLinkToJSON,
} from './InstanceLink';
import type { WarningType } from './WarningType';
import {
    WarningTypeFromJSON,
    WarningTypeFromJSONTyped,
    WarningTypeToJSON,
} from './WarningType';

/**
 * Request to create a Proforma(folio) transactions for a reservation. The request generates the internal transactions required to create a Proforma report. After this request is called , the Opera Proforma report should be called so that these transactions can be used in the report. If the report is called without this request, the report will be incorrect and will not have the complete information. If the report is not used after this request, the internal transactions will be removed either by night audit or by the next request.
 * @export
 * @interface ReservationFolio
 */
export interface ReservationFolio {
    /**
     * 
     * @type {CreateProformaCriteriaType}
     * @memberof ReservationFolio
     */
    criteria?: CreateProformaCriteriaType;
    /**
     * 
     * @type {Array<InstanceLink>}
     * @memberof ReservationFolio
     */
    links?: Array<InstanceLink>;
    /**
     * Used in conjunction with the Success element to define a business error.
     * @type {Array<WarningType>}
     * @memberof ReservationFolio
     */
    warnings?: Array<WarningType>;
}

/**
 * Check if a given object implements the ReservationFolio interface.
 */
export function instanceOfReservationFolio(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function ReservationFolioFromJSON(json: any): ReservationFolio {
    return ReservationFolioFromJSONTyped(json, false);
}

export function ReservationFolioFromJSONTyped(json: any, ignoreDiscriminator: boolean): ReservationFolio {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'criteria': !exists(json, 'criteria') ? undefined : CreateProformaCriteriaTypeFromJSON(json['criteria']),
        'links': !exists(json, 'links') ? undefined : ((json['links'] as Array<any>).map(InstanceLinkFromJSON)),
        'warnings': !exists(json, 'warnings') ? undefined : ((json['warnings'] as Array<any>).map(WarningTypeFromJSON)),
    };
}

export function ReservationFolioToJSON(value?: ReservationFolio | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'criteria': CreateProformaCriteriaTypeToJSON(value.criteria),
        'links': value.links === undefined ? undefined : ((value.links as Array<any>).map(InstanceLinkToJSON)),
        'warnings': value.warnings === undefined ? undefined : ((value.warnings as Array<any>).map(WarningTypeToJSON)),
    };
}

