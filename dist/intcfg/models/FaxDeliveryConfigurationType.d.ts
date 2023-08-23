/**
 * OPERA Cloud Integration Configuration API
 * APIs catering to Integration Configuration in OPERA Cloud.  Operations such as get Hotel Interface Types, or get UDF mappings can be found in this module.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import type { EmailDeliveryFormatType } from './EmailDeliveryFormatType';
import type { EmailDeliveryMethodType } from './EmailDeliveryMethodType';
/**
 * Email content configuration template
 * @export
 * @interface FaxDeliveryConfigurationType
 */
export interface FaxDeliveryConfigurationType {
    /**
     *
     * @type {string}
     * @memberof FaxDeliveryConfigurationType
     */
    fromAddress?: string;
    /**
     *
     * @type {string}
     * @memberof FaxDeliveryConfigurationType
     */
    userId?: string;
    /**
     *
     * @type {string}
     * @memberof FaxDeliveryConfigurationType
     */
    userPassword?: string;
    /**
     *
     * @type {string}
     * @memberof FaxDeliveryConfigurationType
     */
    serverName?: string;
    /**
     *
     * @type {string}
     * @memberof FaxDeliveryConfigurationType
     */
    subject?: string;
    /**
     *
     * @type {string}
     * @memberof FaxDeliveryConfigurationType
     */
    body?: string;
    /**
     *
     * @type {EmailDeliveryFormatType}
     * @memberof FaxDeliveryConfigurationType
     */
    format?: EmailDeliveryFormatType;
    /**
     *
     * @type {EmailDeliveryMethodType}
     * @memberof FaxDeliveryConfigurationType
     */
    type?: EmailDeliveryMethodType;
    /**
     * Hotel code
     * @type {string}
     * @memberof FaxDeliveryConfigurationType
     */
    hotelId?: string;
    /**
     *
     * @type {string}
     * @memberof FaxDeliveryConfigurationType
     */
    faxPrefix?: string;
    /**
     *
     * @type {string}
     * @memberof FaxDeliveryConfigurationType
     */
    faxSuffix?: string;
    /**
     *
     * @type {string}
     * @memberof FaxDeliveryConfigurationType
     */
    faxDomain?: string;
    /**
     * This option must be selected in order to activate faxing from an SMTP compliant fax server
     * @type {boolean}
     * @memberof FaxDeliveryConfigurationType
     */
    activateEmailToFax?: boolean;
}
/**
 * Check if a given object implements the FaxDeliveryConfigurationType interface.
 */
export declare function instanceOfFaxDeliveryConfigurationType(value: object): boolean;
export declare function FaxDeliveryConfigurationTypeFromJSON(json: any): FaxDeliveryConfigurationType;
export declare function FaxDeliveryConfigurationTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): FaxDeliveryConfigurationType;
export declare function FaxDeliveryConfigurationTypeToJSON(value?: FaxDeliveryConfigurationType | null): any;
