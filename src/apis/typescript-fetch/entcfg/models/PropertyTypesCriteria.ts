/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Enterprise Configuration API
 * APIs to cater for Enterprise Configuration functionality in OPERA Cloud. <br /><br In this module, you can configure a variety of options related to your properties such as their locations, facilities, and local attractions. The available options are dependant on the active controls at your property.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
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
import type { PropertyTypeType } from './PropertyTypeType';
import {
    PropertyTypeTypeFromJSON,
    PropertyTypeTypeFromJSONTyped,
    PropertyTypeTypeToJSON,
} from './PropertyTypeType';
import type { WarningType } from './WarningType';
import {
    WarningTypeFromJSON,
    WarningTypeFromJSONTyped,
    WarningTypeToJSON,
} from './WarningType';

/**
 * Request object for creating Property Types.
 * @export
 * @interface PropertyTypesCriteria
 */
export interface PropertyTypesCriteria {
    /**
     * List of Property Types.
     * @type {Array<PropertyTypeType>}
     * @memberof PropertyTypesCriteria
     */
    propertyTypes?: Array<PropertyTypeType>;
    /**
     * 
     * @type {Array<InstanceLink>}
     * @memberof PropertyTypesCriteria
     */
    links?: Array<InstanceLink>;
    /**
     * Used in conjunction with the Success element to define a business error.
     * @type {Array<WarningType>}
     * @memberof PropertyTypesCriteria
     */
    warnings?: Array<WarningType>;
}

/**
 * Check if a given object implements the PropertyTypesCriteria interface.
 */
export function instanceOfPropertyTypesCriteria(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function PropertyTypesCriteriaFromJSON(json: any): PropertyTypesCriteria {
    return PropertyTypesCriteriaFromJSONTyped(json, false);
}

export function PropertyTypesCriteriaFromJSONTyped(json: any, ignoreDiscriminator: boolean): PropertyTypesCriteria {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'propertyTypes': !exists(json, 'propertyTypes') ? undefined : ((json['propertyTypes'] as Array<any>).map(PropertyTypeTypeFromJSON)),
        'links': !exists(json, 'links') ? undefined : ((json['links'] as Array<any>).map(InstanceLinkFromJSON)),
        'warnings': !exists(json, 'warnings') ? undefined : ((json['warnings'] as Array<any>).map(WarningTypeFromJSON)),
    };
}

export function PropertyTypesCriteriaToJSON(value?: PropertyTypesCriteria | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'propertyTypes': value.propertyTypes === undefined ? undefined : ((value.propertyTypes as Array<any>).map(PropertyTypeTypeToJSON)),
        'links': value.links === undefined ? undefined : ((value.links as Array<any>).map(InstanceLinkToJSON)),
        'warnings': value.warnings === undefined ? undefined : ((value.warnings as Array<any>).map(WarningTypeToJSON)),
    };
}

