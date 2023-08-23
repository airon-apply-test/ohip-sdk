/**
 * OPERA Cloud Rate API
 * APIs to cater for Rate Availability functionality in OPERA Cloud. <br /><br /> Rate Management provides all the tools you need to effectively define and manage the rate structures for a property in OPERA Cloud. Some of the things you can do include creating and managing rate codes, rate classes, rate categories, display sets, rate strategies, as well as managing promotion groups and codes.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import type { NegRateAccessType } from './NegRateAccessType';
import type { ProfileNameType } from './ProfileNameType';
import type { ProfileTypeType } from './ProfileTypeType';
import type { UniqueIDType } from './UniqueIDType';
/**
 * Profile details
 * @export
 * @interface NegotiatedProfileType
 */
export interface NegotiatedProfileType {
    /**
     * Unique Id that references an object uniquely in the system.
     * @type {Array<UniqueIDType>}
     * @memberof NegotiatedProfileType
     */
    profileIdList?: Array<UniqueIDType>;
    /**
     *
     * @type {ProfileNameType}
     * @memberof NegotiatedProfileType
     */
    profileName?: ProfileNameType;
    /**
     * collection of Negotiated rates.
     * @type {Array<NegRateAccessType>}
     * @memberof NegotiatedProfileType
     */
    rateInfoList?: Array<NegRateAccessType>;
    /**
     *
     * @type {ProfileTypeType}
     * @memberof NegotiatedProfileType
     */
    profileType?: ProfileTypeType;
}
/**
 * Check if a given object implements the NegotiatedProfileType interface.
 */
export declare function instanceOfNegotiatedProfileType(value: object): boolean;
export declare function NegotiatedProfileTypeFromJSON(json: any): NegotiatedProfileType;
export declare function NegotiatedProfileTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): NegotiatedProfileType;
export declare function NegotiatedProfileTypeToJSON(value?: NegotiatedProfileType | null): any;
