/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Price Availability Rate Async API
 * APIs to cater for Price and Rate Availability Asynchronous functionality in OPERA Cloud.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import type { RestrictionStatusesType } from './RestrictionStatusesType';
import {
    RestrictionStatusesTypeFromJSON,
    RestrictionStatusesTypeFromJSONTyped,
    RestrictionStatusesTypeToJSON,
} from './RestrictionStatusesType';

/**
 * The RestrictionStatusType is used to indicate the type of restriction applied. An enumerated restriction type is defined in the attribute group.
 * @export
 * @interface RestrictionStatusType
 */
export interface RestrictionStatusType {
    /**
     * 
     * @type {RestrictionStatusesType}
     * @memberof RestrictionStatusType
     */
    code?: RestrictionStatusesType;
    /**
     * Defines restriction in conjunction with Restriction Code. Value must be specified for Restriction Codes MinimumStayThrough, MaximumStayThrough, MinimumLengthOfStay, MaximumLengthOfStay, MinimumAdvanceBooking, MaximumAdvanceBooking.
     * @type {number}
     * @memberof RestrictionStatusType
     */
    unit?: number;
    /**
     * Indicates Length of Stay 1 is not available. Defines restriction in conjunction with Restriction Code LengthOfStayNotAvailable.
     * @type {boolean}
     * @memberof RestrictionStatusType
     */
    lengthOfStay1?: boolean;
    /**
     * Indicates Length of Stay 2 is not available. Defines restriction in conjunction with Restriction Code LengthOfStayNotAvailable.
     * @type {boolean}
     * @memberof RestrictionStatusType
     */
    lengthOfStay2?: boolean;
    /**
     * Indicates Length of Stay 3 is not available. Defines restriction in conjunction with Restriction Code LengthOfStayNotAvailable.
     * @type {boolean}
     * @memberof RestrictionStatusType
     */
    lengthOfStay3?: boolean;
    /**
     * Indicates Length of Stay 4 is not available. Defines restriction in conjunction with Restriction Code LengthOfStayNotAvailable.
     * @type {boolean}
     * @memberof RestrictionStatusType
     */
    lengthOfStay4?: boolean;
    /**
     * Indicates Length of Stay 5 is not available. Defines restriction in conjunction with Restriction Code LengthOfStayNotAvailable.
     * @type {boolean}
     * @memberof RestrictionStatusType
     */
    lengthOfStay5?: boolean;
    /**
     * Indicates Length of Stay 6 is not available. Defines restriction in conjunction with Restriction Code LengthOfStayNotAvailable.
     * @type {boolean}
     * @memberof RestrictionStatusType
     */
    lengthOfStay6?: boolean;
    /**
     * Indicates Length of Stay 7 or more is not available. Defines restriction in conjunction with Restriction Code LengthOfStayNotAvailable.
     * @type {boolean}
     * @memberof RestrictionStatusType
     */
    lengthOfStay7?: boolean;
}

/**
 * Check if a given object implements the RestrictionStatusType interface.
 */
export function instanceOfRestrictionStatusType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function RestrictionStatusTypeFromJSON(json: any): RestrictionStatusType {
    return RestrictionStatusTypeFromJSONTyped(json, false);
}

export function RestrictionStatusTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): RestrictionStatusType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'code': !exists(json, 'code') ? undefined : RestrictionStatusesTypeFromJSON(json['code']),
        'unit': !exists(json, 'unit') ? undefined : json['unit'],
        'lengthOfStay1': !exists(json, 'lengthOfStay1') ? undefined : json['lengthOfStay1'],
        'lengthOfStay2': !exists(json, 'lengthOfStay2') ? undefined : json['lengthOfStay2'],
        'lengthOfStay3': !exists(json, 'lengthOfStay3') ? undefined : json['lengthOfStay3'],
        'lengthOfStay4': !exists(json, 'lengthOfStay4') ? undefined : json['lengthOfStay4'],
        'lengthOfStay5': !exists(json, 'lengthOfStay5') ? undefined : json['lengthOfStay5'],
        'lengthOfStay6': !exists(json, 'lengthOfStay6') ? undefined : json['lengthOfStay6'],
        'lengthOfStay7': !exists(json, 'lengthOfStay7') ? undefined : json['lengthOfStay7'],
    };
}

export function RestrictionStatusTypeToJSON(value?: RestrictionStatusType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'code': RestrictionStatusesTypeToJSON(value.code),
        'unit': value.unit,
        'lengthOfStay1': value.lengthOfStay1,
        'lengthOfStay2': value.lengthOfStay2,
        'lengthOfStay3': value.lengthOfStay3,
        'lengthOfStay4': value.lengthOfStay4,
        'lengthOfStay5': value.lengthOfStay5,
        'lengthOfStay6': value.lengthOfStay6,
        'lengthOfStay7': value.lengthOfStay7,
    };
}

